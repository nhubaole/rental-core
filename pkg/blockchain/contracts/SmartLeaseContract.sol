// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SmartLeaseContract {
    enum PreRentalStatus { Activated, Signed, PaidDeposit }
    enum RentalProcessStatus { FirstPhase, RecordedMeter, Unpaid, Paid }
    enum PostRentalStatus { NotRequested, UnreturnedDeposit, Completed }

    struct RentalAgreementInfo {
        uint256 roomId;
        uint256 actualPrice;
        uint256 depositAmount;
        uint256 beginDate;
        uint256 endDate;
        string paymentMethod;
        string electricityMethod;
        uint256 electricityCost;
        string waterMethod;
        uint256 waterCost;
        uint256 internetCost;
        uint256 parkingFee;
        string responsibilityA;
        string responsibilityB;
        string generalResponsibility;
        string signatureA;
        uint256 signedTimeA;
        string signatureB;
        uint256 signedTimeB;
        uint256 createdAt;
        uint256 updatedAt;
        uint256 deletedAt;
        uint256 contractTemplateId;
    }

    struct RentalAgreement {
        uint256 contractId;
        string contractCode;
        address landlord;
        address tenant;
        PreRentalStatus preRentalStatus;
        RentalProcessStatus rentalProcessStatus;
        PostRentalStatus postRentalStatus;
        uint256 status;
    }

    struct BillingRecord {
        uint256 billingId;
        string billingCode;
        uint256 contractId;
        uint256 additionFee;
        string additionNote;
        uint256 totalAmount;
        uint256 month;
        uint256 year;
        uint256 paidTime;
        uint256 createdAt;
        uint256 updatedAt;
        uint256 deletedAt;
        uint256 oldWaterIndex;
        uint256 oldElectricityIndex;
        uint256 newWaterIndex;
        uint256 newElectricityIndex;
        uint256 totalWaterCost;
        uint256 totalElectricityCost;
        uint256 status;
    }

    struct ReturnRequest {
        uint256 requestId;
        uint256 contractId;
        bytes32 reason;
        uint256 returnDate;
        uint256 status;
        uint256 deductAmount;
        uint256 totalReturnDeposit;
        address createdUser;
        uint256 createdAt;
        uint256 updatedAt;
        uint256 deletedAt;
    }

    RentalAgreement public agreement;
    RentalAgreementInfo public agreementInfo;
    BillingRecord[] public billingRecords;
    ReturnRequest public returnRequest;

    uint256 private nextRequestId = 1;

    modifier onlyLandlord() {
        require(msg.sender == agreement.landlord, "Only landlord can perform this action");
        _;
    }

    modifier onlyTenant() {
        require(msg.sender == agreement.tenant, "Only tenant can perform this action");
        _;
    }

    constructor(
        string memory _contractCode,
        address _landlord,
        address _tenant,
        RentalAgreementInfo memory _agreementInfo
    ) {
        agreement.contractCode = _contractCode;
        agreement.landlord = _landlord;
        agreement.tenant = _tenant;
        agreement.preRentalStatus = PreRentalStatus.Activated;
        agreement.rentalProcessStatus = RentalProcessStatus.FirstPhase;
        agreement.postRentalStatus = PostRentalStatus.NotRequested;
        agreementInfo = _agreementInfo;
    }

    function getContractDetails() external view returns (
        address landlord,
        address tenant,
        uint256 roomId,
        uint256 actualPrice,
        uint256 depositAmount,
        uint256 beginDate,
        uint256 endDate,
        string memory contractCode
    ) {
        return (
            agreement.landlord,
            agreement.tenant,
            agreementInfo.roomId,
            agreementInfo.actualPrice,
            agreementInfo.depositAmount,
            agreementInfo.beginDate,
            agreementInfo.endDate,
            agreement.contractCode
        );
    }

    // Pre-Rental Stage Functions
    function signContractByTenant(string memory _signatureB) public onlyTenant {
        agreementInfo.signatureB = _signatureB;
        agreementInfo.signedTimeB = block.timestamp;
        agreement.preRentalStatus = PreRentalStatus.Signed;
    }
    
    function payDeposit() public payable onlyTenant {
        require(agreement.preRentalStatus == PreRentalStatus.Signed, "Contract must be signed before paying deposit");
        require(msg.value == agreementInfo.depositAmount, "Deposit amount incorrect");
        agreement.preRentalStatus = PreRentalStatus.PaidDeposit;
    }

    // Rental Process Stage Functions
    function inputMeterReading() public onlyLandlord {
        require(agreement.preRentalStatus == PreRentalStatus.PaidDeposit, "Deposit must be paid before recording meter");
        require(agreement.rentalProcessStatus != RentalProcessStatus.FirstPhase, "Meter reading can only be recorded after the first phase");

        agreement.rentalProcessStatus = RentalProcessStatus.RecordedMeter;
    }

    function createBill(
        uint256 _totalAmount,
        uint256 _month,
        uint256 _year,
        uint256 _additionFee,
        string memory _additionNote,
        uint256 _oldWaterIndex,
        uint256 _newWaterIndex,
        uint256 _oldElectricityIndex,
        uint256 _newElectricityIndex
    ) public onlyLandlord {
        require(
            agreement.rentalProcessStatus == RentalProcessStatus.FirstPhase || 
            agreement.rentalProcessStatus == RentalProcessStatus.RecordedMeter, 
            "Bill creation requires first phase or recorded meter reading in subsequent phases"
        );

        uint256 waterCost = (_newWaterIndex - _oldWaterIndex) * agreementInfo.waterCost;
        uint256 electricityCost = (_newElectricityIndex - _oldElectricityIndex) * agreementInfo.electricityCost;

        billingRecords.push(BillingRecord({
            billingId: billingRecords.length + 1,
            billingCode: string(abi.encodePacked(agreement.contractCode, "-", uint2str(_month), "-", uint2str(_year))),
            contractId: agreement.contractId,
            additionFee: _additionFee,
            additionNote: _additionNote,
            totalAmount: _totalAmount + waterCost + electricityCost,
            month: _month,
            year: _year,
            paidTime: 0,
            createdAt: block.timestamp,
            updatedAt: block.timestamp,
            deletedAt: 0,
            oldWaterIndex: _oldWaterIndex,
            oldElectricityIndex: _oldElectricityIndex,
            newWaterIndex: _newWaterIndex,
            newElectricityIndex: _newElectricityIndex,
            totalWaterCost: waterCost,
            totalElectricityCost: electricityCost,
            status: 0
        }));

        agreement.rentalProcessStatus = RentalProcessStatus.Unpaid;
    }


    function payBill(uint256 _billingId) public payable onlyTenant {
        require(agreement.rentalProcessStatus == RentalProcessStatus.Unpaid, "No unpaid bill to pay");
        require(msg.value == billingRecords[_billingId - 1].totalAmount, "Bill amount incorrect");
        billingRecords[_billingId - 1].paidTime = block.timestamp;
        agreement.rentalProcessStatus = RentalProcessStatus.Paid;
    }

    // Post-Rental Stage Functions
    function createReturnRequest(
        uint256 _deductAmount,
        uint256 _totalReturnDeposit,
        bytes32 _reason
    ) public onlyTenant {
        require(agreement.postRentalStatus == PostRentalStatus.NotRequested, "Return request already exists or completed");

        returnRequest = ReturnRequest({
            requestId: nextRequestId,
            contractId: agreement.contractId,
            reason: _reason,
            returnDate: block.timestamp,
            status: 0,
            deductAmount: _deductAmount,
            totalReturnDeposit: _totalReturnDeposit,
            createdUser: msg.sender,
            createdAt: block.timestamp,
            updatedAt: block.timestamp,
            deletedAt: 0
        });

        nextRequestId++; // Increment the counter for the next request
        agreement.postRentalStatus = PostRentalStatus.UnreturnedDeposit;
    }

    function approveReturnRequest() public onlyLandlord {
        require(agreement.postRentalStatus == PostRentalStatus.UnreturnedDeposit, "No pending return request to approve");
        payable(agreement.tenant).transfer(returnRequest.totalReturnDeposit);
        returnRequest.status = 1;
        agreement.postRentalStatus = PostRentalStatus.Completed;
    }

    // Helper function to convert uint to string
    function uint2str(uint _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint j = _i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len;
        while (_i != 0) {
            k = k - 1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }
}
