// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ContractManagement {
    enum PreRentalStatus { Activated, Signed, PaidDeposit }
    enum RentalProcessStatus { FirstPhase, RecordedMeter, Unpaid, Paid }
    enum PostRentalStatus { NotRequested, UnreturnedDeposit, Completed }
    
    struct MContract {
        uint id;
        string code;
        uint landlord;
        uint tenant;
        uint roomId;
        uint256 actualPrice;
        uint256 deposit;
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
        uint contractTemplateId;
        PreRentalStatus preRentalStatus;
        RentalProcessStatus rentalProcessStatus;
        PostRentalStatus postRentalStatus;
        uint256 createdAt;
        uint256 updatedAt;
        bool exists;
    }

    mapping(uint => MContract) public mContracts;

    event MContractCreated(
        uint id,
        string code,
        uint landlord,
        uint tenant,
        uint roomId,
        uint256 actualPrice,
        uint256 deposit,
        uint256 beginDate,
        uint256 endDate
    );

    event MContractUpdated(uint id);
    event MContractDeleted(uint id);

    // Tạo hợp đồng mới
    function createMContract(
        uint _id,
        string memory _code,
        uint _landlord,
        uint _tenant,
        uint _roomId,
        uint256 _actualPrice,
        uint256 _deposit,
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
        string memory _signatureB,
        uint256 _signedTimeB,
        uint _contractTemplateId
    ) public {
        MContract storage mContract = mContracts[_id];
        require(mContract.exists == false, "MContract already exists");

        mContracts[_id] = MContract({
            id: _id,
            code: _code,
            landlord: _landlord,
            tenant: _tenant,
            roomId: _roomId,
            actualPrice: _actualPrice,
            deposit: _deposit,
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
            signatureB: _signatureB,
            signedTimeB: _signedTimeB,
            contractTemplateId: _contractTemplateId,
            preRentalStatus: PreRentalStatus.Activated,
            rentalProcessStatus: RentalProcessStatus.FirstPhase,
            postRentalStatus: PostRentalStatus.NotRequested,
            createdAt: block.timestamp,
            updatedAt: block.timestamp,
            exists: true
        });

        emit MContractCreated(
            _id,
            _code,
            _landlord,
            _tenant,
            _roomId,
            _actualPrice,
            _deposit,
            _beginDate,
            _endDate
        );
    }

    // Lấy thông tin hợp đồng
    function getMContract(
        uint _id
    )
        public
        view
        returns (
            uint,
            string memory,
            uint,
            uint,
            uint,
            uint256,
            uint256,
            uint256,
            uint256,
            string memory,
            string memory,
            uint256,
            string memory,
            uint256,
            uint256,
            uint256,
            string memory,
            string memory,
            string memory,
            string memory,
            uint256,
            string memory,
            uint256,
            uint,
            PreRentalStatus,
            RentalProcessStatus,
            PostRentalStatus,
            uint256,
            uint256
        )
    {
        MContract storage mContract = mContracts[_id];
        require(mContract.exists == true, "MContract does not exist");

        return (
            mContract.id,
            mContract.code,
            mContract.landlord,
            mContract.tenant,
            mContract.roomId,
            mContract.actualPrice,
            mContract.deposit,
            mContract.beginDate,
            mContract.endDate,
            mContract.paymentMethod,
            mContract.electricityMethod,
            mContract.electricityCost,
            mContract.waterMethod,
            mContract.waterCost,
            mContract.internetCost,
            mContract.parkingFee,
            mContract.responsibilityA,
            mContract.responsibilityB,
            mContract.generalResponsibility,
            mContract.signatureA,
            mContract.signedTimeA,
            mContract.signatureB,
            mContract.signedTimeB,
            mContract.contractTemplateId,
            mContract.preRentalStatus,
            mContract.rentalProcessStatus,
            mContract.postRentalStatus,
            mContract.createdAt,
            mContract.updatedAt
        );
    }
}
