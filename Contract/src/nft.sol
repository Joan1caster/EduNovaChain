// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract EducationInnovationNFT is ERC721URIStorage, Ownable {
    uint256 private _tokenIds;

    struct Innovation {
        string tokenID;
        string metadata;
        string IPFSCID;
        uint256 price;
        address creator;
        address owner;
        bool isForSale;
        uint8 royaltyPercentage;
        uint256 creationTime;
        uint256 lastSaleTime;
    }

    mapping(uint256 => Innovation) public innovations;
    mapping(address => uint256[]) public creatorInnovations;
    mapping(string => bool) public ExistsCID;

    event InnovationCreate(
        address owner,
        uint256 tolenId,
        uint256 creationTime
    );

    event InnovationPurchased(
        uint256 indexed tokenId,
        address indexed seller,
        address indexed buyer
    );
    event InnovationPriceUpdated(uint256 indexed tokenId, uint256 newPrice);

    constructor()
        ERC721("EducationInnovationNFT", "EINFT")
        Ownable(msg.sender)
    {}

    function createInnovation(
        string calldata tokenID,
        string calldata metadata,
        string calldata IPFSCID,
        uint256 price,
        bool isForSale
    ) public returns (uint256) {
        require(!ExistsCID[IPFSCID], "repeated IPFSCID");
        _tokenIds += 1;
        uint256 newItemId = _tokenIds;

        _safeMint(msg.sender, newItemId);
        _setTokenURI(newItemId, IPFSCID);

        innovations[newItemId] = Innovation(
            tokenID,
            metadata,
            IPFSCID,
            price,
            msg.sender,
            msg.sender,
            isForSale,
            0,
            block.timestamp,
            0
        );

        creatorInnovations[msg.sender].push(newItemId);

        emit InnovationCreate(msg.sender, newItemId, block.timestamp);

        return newItemId;
    }

    function purchaseInnovation(uint256 tokenId) public payable {
        require(
            innovations[tokenId].creator != address(0),
            "Innovation does not exist"
        );
        require(innovations[tokenId].isForSale, "Innovation is not for sale");
        require(
            msg.value >= innovations[tokenId].price,
            "Insufficient payment"
        );

        address seller = ownerOf(tokenId);
        address buyer = msg.sender;

        _transfer(seller, buyer, tokenId);
        payable(seller).transfer(msg.value);

        innovations[tokenId].isForSale = false;

        emit InnovationPurchased(tokenId, innovations[tokenId].owner, buyer);
    }

    function updateInnovationPrice(uint256 tokenId, uint256 newPrice) public {
        require(ownerOf(tokenId) == msg.sender, "Not the owner");
        innovations[tokenId].price = newPrice;
        innovations[tokenId].isForSale = true;

        emit InnovationPriceUpdated(tokenId, newPrice);
    }

    function getInnovationsByCreator(
        address creator
    ) public view returns (uint256[] memory) {
        return creatorInnovations[creator];
    }

    function getInnovationDetails(
        uint256 tokenId
    ) public view returns (Innovation memory) {
        require(
            innovations[tokenId].creator != address(0),
            "Innovation does not exist"
        );
        return innovations[tokenId];
    }

    function updataSale(uint256 tokenId) public {
        require(
            innovations[tokenId].creator != address(0),
            "Innovation does not exist"
        );
        innovations[tokenId].isForSale = true;
    }
}
