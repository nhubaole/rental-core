// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RoomManagement {
    
    struct Room {
        uint id;
        uint owner;
        uint status;
        bool isRent;
        uint256 totalPrice;
        uint256 deposit;
        uint256 createdAt;
        uint256 updatedAt;
        bool exists; 
    }
    
    mapping(uint => Room) public rooms;
    
    event RoomCreated(uint id, uint owner, uint256 totalPrice, uint256 deposit, uint status, bool isRent);
    event RoomUpdated(uint id, uint owner, uint256 totalPrice, uint256 deposit, uint status, bool isRent);
    event RoomDeleted(uint id);

    // Tạo phòng mới
    function createRoom(uint _id, uint _owner, uint256 _totalPrice, uint256 _deposit, uint _status, bool _isRent) public {
        Room storage room = rooms[_id];
        require(room.exists == false, "Room already exists");
        
        rooms[_id] = Room({
            id: _id,
            owner: _owner,
            status: _status,
            isRent: _isRent,
            totalPrice: _totalPrice,
            deposit: _deposit,
            createdAt: block.timestamp,
            updatedAt: block.timestamp,
            exists: true
        });
        
        emit RoomCreated(_id, _owner, _totalPrice, _deposit, _status, _isRent);
    }

    // Cập nhật thông tin phòng
    function updateRoom(uint _id, uint256 _totalPrice, uint256 _deposit, uint _status, bool _isRent) public {
        Room storage room = rooms[_id];
        require(room.exists == true, "Room does not exist");
        
        room.totalPrice = _totalPrice;
        room.deposit = _deposit;
        room.status = _status;
        room.isRent = _isRent;
        room.updatedAt = block.timestamp;
        
        emit RoomUpdated(_id, room.owner, _totalPrice, _deposit, _status, _isRent);
    }

    // Xóa phòng
    function deleteRoom(uint _id) public {
        Room storage room = rooms[_id];
        require(room.exists == true, "Room does not exist");
        
        room.exists = false; // Đánh dấu phòng là không tồn tại nữa
        
        emit RoomDeleted(_id);
    }

    // Lấy thông tin phòng
    function getRoom(uint _id) public view returns (uint, uint, uint256, uint256, uint, bool, uint256, uint256) {
        Room storage room = rooms[_id];
        require(room.exists == true, "Room does not exist");
        
        return (
            room.id,
            room.owner,
            room.totalPrice,
            room.deposit,
            room.status,
            room.isRent,
            room.createdAt,
            room.updatedAt
        );
    }
}
