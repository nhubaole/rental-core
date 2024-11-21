// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./SmartLeaseContract.sol";  // Ensure SmartLeaseContract is available in the same directory

contract LeaseAgreementProducerContract {
    // Mapping to store lease contracts associated with each participant (landlord or tenant)
    mapping(address => address[]) public contractsByParticipant;

    // Event emitted when a new lease contract is created
    event LeaseContractCreated(
        address indexed contractAddress,
        address indexed landlord,
        address indexed tenant,
        uint256 roomId,
        uint256 actualPrice,
        uint256 depositAmount,
        uint256 beginDate,
        uint256 endDate
    );

    // Function to create and initialize a new SmartLeaseContract with Party A's signature
function createLeaseContract(
        address tenant,
        uint256 roomId,
        uint256 actualPrice,
        uint256 depositAmount,
        uint256 beginDate,
        uint256 endDate,
        string memory contractCode,
        string memory signatureA,        // Landlord's signature
        uint256 signedTimeA,       // Landlord's signing timestamp
        string memory paymentMethod,
        string memory electricityMethod,
        uint256 electricityCost,
        string memory waterMethod,
        uint256 waterCost,
        uint256 internetCost,
        uint256 parkingFee,
        string memory responsibilityA,
        string memory responsibilityB,
        string memory generalResponsibility,
        uint256 contractTemplateId
    ) public {
        // Initialize the RentalAgreementInfo struct with additional fields
        SmartLeaseContract.RentalAgreementInfo memory agreementInfo = SmartLeaseContract.RentalAgreementInfo({
            roomId: roomId,
            actualPrice: actualPrice,
            depositAmount: depositAmount,
            beginDate: beginDate,
            endDate: endDate,
            paymentMethod: paymentMethod,
            electricityMethod: electricityMethod,
            electricityCost: electricityCost,
            waterMethod: waterMethod,
            waterCost: waterCost,
            internetCost: internetCost,
            parkingFee: parkingFee,
            responsibilityA: responsibilityA,
            responsibilityB: responsibilityB,
            generalResponsibility: generalResponsibility,
            signatureA: signatureA,
            signedTimeA: signedTimeA,
            signatureB: "",
            signedTimeB: 0,
            createdAt: block.timestamp,
            updatedAt: block.timestamp,
            deletedAt: 0,
            contractTemplateId: contractTemplateId
        });

        // Deploy a new SmartLeaseContract with the struct and other required parameters
        SmartLeaseContract leaseContract = new SmartLeaseContract(
            contractCode,
            msg.sender,    // Landlord's address
            tenant,        // Tenant's address
            agreementInfo  // Pass the RentalAgreementInfo struct with full parameters
        );

        // Store the contract address for both landlord and tenant
        contractsByParticipant[msg.sender].push(address(leaseContract));  // Landlord's contracts
        contractsByParticipant[tenant].push(address(leaseContract));      // Tenant's contracts

        // Emit event with contract details
        emit LeaseContractCreated(
            address(leaseContract),
            msg.sender,
            tenant,
            roomId,
            actualPrice,
            depositAmount,
            beginDate,
            endDate
        );
    }

    // Function to retrieve all contract addresses associated with a specific participant
    function getContractsByParticipant(address participant) public view returns (address[] memory) {
        return contractsByParticipant[participant];
    }

    // Function to get details of a specific SmartLeaseContract by its address
    function getLeaseContractDetails(address contractAddress) public view returns (
        address landlord,
        address tenant,
        uint256 roomId,
        uint256 actualPrice,
        uint256 depositAmount,
        uint256 beginDate,
        uint256 endDate,
        string memory contractCode
    ) {
        SmartLeaseContract leaseContract = SmartLeaseContract(contractAddress);
        return leaseContract.getContractDetails();
    }
}
