// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract LeaseContractManagement {
    enum PreRentalStatus { Activated, Signed, PaidDeposit }
    enum RentalProcessStatus { FirstPhase, RecordedMeter, Unpaid, Paid }
    enum PostRentalStatus { NotRequested, UnreturnedDeposit, Completed }

    struct LeaseContract {
        uint256 contractId;
        string contractCode;
        address landlord;
        address tenant;
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
        PreRentalStatus preRentalStatus;
        RentalProcessStatus rentalProcessStatus;
        PostRentalStatus postRentalStatus;
        uint256 status; // 0: Pending, 1: Active, 2: Terminated
        bool exists;
    }

    mapping(uint256 => LeaseContract) public leaseContracts;
    uint256 public nextContractId = 1;

    event LeaseContractCreated(uint256 contractId, string contractCode, address landlord, address tenant, uint256 roomId);
    event BillingRecordCreated(uint256 contractId, string billingCode);
    event ReturnRequestCreated(uint256 contractId, string returnCode);

    modifier onlyLandlord(uint256 _contractId) {
        require(msg.sender == leaseContracts[_contractId].landlord, "Only landlord can perform this action");
        _;
    }

    modifier onlyTenant(uint256 _contractId) {
        require(msg.sender == leaseContracts[_contractId].tenant, "Only tenant can perform this action");
        _;
    }

    function createLeaseContract(
        string memory _code,
        address _landlord,
        address _tenant,
        uint256 _roomId,
        uint256 _actualPrice,
        uint256 _depositAmount,
        uint256 _beginDate,
        uint256 _endDate,
        string memory _paymentMethod,
        string memory _electricityMethod,
        uint256 _electricityCost,
        string memory _waterMethod,
        uint256 _waterCost,
        uint256 _internetCost,
        uint256 _parkingFee,
        string memory _responsibilityA,
        string memory _responsibilityB,
        string memory _generalResponsibility,
        string memory _signatureA,
        uint256 _signedTimeA,
        uint256 _contractTemplateId
    ) public {
        require(_landlord != address(0), "Invalid landlord address");
        require(_tenant != address(0), "Invalid tenant address");
        require(_beginDate < _endDate, "Begin date must be before end date");
        require(_depositAmount >= 0, "Deposit must be non-negative");
        require(bytes(_code).length > 0, "Contract code cannot be empty");

        leaseContracts[nextContractId] = LeaseContract({
            contractId: nextContractId,
            contractCode: _code,
            landlord: _landlord,
            tenant: _tenant,
            roomId: _roomId,
            actualPrice: _actualPrice,
            depositAmount: _depositAmount,
            beginDate: _beginDate,
            endDate: _endDate,
            paymentMethod: _paymentMethod,
            electricityMethod: _electricityMethod,
            electricityCost: _electricityCost,
            waterMethod: _waterMethod,
            waterCost: _waterCost,
            internetCost: _internetCost,
            parkingFee: _parkingFee,
            responsibilityA: _responsibilityA,
            responsibilityB: _responsibilityB,
            generalResponsibility: _generalResponsibility,
            signatureA: _signatureA,
            signedTimeA: _signedTimeA,
            signatureB: "",
            signedTimeB: 0,
            createdAt: block.timestamp,
            updatedAt: block.timestamp,
            deletedAt: 0,
            contractTemplateId: _contractTemplateId,
            preRentalStatus: PreRentalStatus.Activated,
            rentalProcessStatus: RentalProcessStatus.FirstPhase,
            postRentalStatus: PostRentalStatus.NotRequested,
            status: 0,
            exists: true
        });

        emit LeaseContractCreated(nextContractId, _code, _landlord, _tenant, _roomId);
        nextContractId++;
    }

    function getContractById(uint256 _contractId) public view returns (LeaseContract memory) {
        require(leaseContracts[_contractId].exists, "Contract does not exist");
        return leaseContracts[_contractId];
    }

    function listContractsByStatus(uint256 _status) public view returns (LeaseContract[] memory) {
        uint256 count = 0;
        for (uint256 i = 1; i < nextContractId; i++) {
            if (leaseContracts[i].status == _status && leaseContracts[i].exists) {
                count++;
            }
        }

        LeaseContract[] memory filteredContracts = new LeaseContract[](count);
        uint256 index = 0;
        for (uint256 i = 1; i < nextContractId; i++) {
            if (leaseContracts[i].status == _status && leaseContracts[i].exists) {
                filteredContracts[index] = leaseContracts[i];
                index++;
            }
        }

        return filteredContracts;
    }

    function signContractByTenant(uint256 _contractId, string memory _signatureB) public onlyTenant(_contractId) {
        LeaseContract storage leaseContract = leaseContracts[_contractId];
        require(leaseContract.preRentalStatus == PreRentalStatus.Activated, "Contract is not ready for signing");

        leaseContract.signatureB = _signatureB;
        leaseContract.signedTimeB = block.timestamp;
        leaseContract.preRentalStatus = PreRentalStatus.Signed;
        leaseContract.updatedAt = block.timestamp;
    }

    function payDeposit(uint256 _contractId) public payable onlyTenant(_contractId) {
        LeaseContract storage leaseContract = leaseContracts[_contractId];
        require(leaseContract.preRentalStatus == PreRentalStatus.Signed, "Contract must be signed before paying deposit");
        require(msg.value == leaseContract.depositAmount, "Deposit amount incorrect");

        leaseContract.preRentalStatus = PreRentalStatus.PaidDeposit;
        leaseContract.updatedAt = block.timestamp;
    }

    function inputMeterReading(uint256 _contractId) public onlyLandlord(_contractId) {
        LeaseContract storage leaseContract = leaseContracts[_contractId];
        require(leaseContract.preRentalStatus == PreRentalStatus.PaidDeposit, "Deposit must be paid before recording meter");
        require(leaseContract.rentalProcessStatus != RentalProcessStatus.FirstPhase, "Meter reading can only be recorded after the first phase");

        leaseContract.rentalProcessStatus = RentalProcessStatus.RecordedMeter;
        leaseContract.updatedAt = block.timestamp;
    }

    function createBill(
        string memory _billingCode,
        uint256 _contractId
    ) public onlyLandlord(_contractId) {
        LeaseContract storage leaseContract = leaseContracts[_contractId];
        require(
            leaseContract.rentalProcessStatus == RentalProcessStatus.FirstPhase ||
            leaseContract.rentalProcessStatus == RentalProcessStatus.RecordedMeter,
            "Bill creation requires first phase or recorded meter reading in subsequent phases"
        );

        leaseContract.rentalProcessStatus = RentalProcessStatus.Unpaid;
        emit BillingRecordCreated(_contractId, _billingCode);
    }

    function payBill(uint256 _contractId) public payable onlyTenant(_contractId) {
        LeaseContract storage leaseContract = leaseContracts[_contractId];
        require(leaseContract.rentalProcessStatus == RentalProcessStatus.Unpaid, "No unpaid bill to pay");

        leaseContract.rentalProcessStatus = RentalProcessStatus.Paid;
    }

    function createReturnRequest(uint256 _contractId, string memory _returnCode) public onlyTenant(_contractId) {
        LeaseContract storage leaseContract = leaseContracts[_contractId];
        require(leaseContract.postRentalStatus == PostRentalStatus.NotRequested, "Return request already exists");

        leaseContract.postRentalStatus = PostRentalStatus.UnreturnedDeposit;
        emit ReturnRequestCreated(_contractId, _returnCode);
    }

    function approveReturnRequest(uint256 _contractId) public onlyLandlord(_contractId) {
        LeaseContract storage leaseContract = leaseContracts[_contractId];
        require(leaseContract.postRentalStatus == PostRentalStatus.UnreturnedDeposit, "No pending return request");

        leaseContract.postRentalStatus = PostRentalStatus.Completed;
    }

    function uint2str(uint256 _i) internal pure returns (string memory) {
        if (_i == 0) return "0";
        uint256 j = _i;
        uint256 length;
        while (j != 0) {
            length++;
            j /= 10;
        }
        bytes memory bstr = new bytes(length);
        while (_i != 0) {
            bstr[--length] = bytes1(uint8(48 + _i % 10));
            _i /= 10;
        }
        return string(bstr);
    }
}
