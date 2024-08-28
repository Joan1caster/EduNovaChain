// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract EducationInnovationNFT is ERC721URIStorage, Ownable {
    uint256 private _tokenIds;

    struct Innovation {
        string title;
        string description;
        uint256 price;
        address creator;
        bool isForSale;
    }

    mapping(uint256 => Innovation) public innovations;
    mapping(address => uint256[]) public creatorInnovations;

    event InnovationCreated(
        uint256 indexed tokenId,
        address indexed creator,
        string title
    );
    event InnovationPurchased(
        uint256 indexed tokenId,
        address indexed buyer,
        uint256 price
    );
    event InnovationPriceUpdated(uint256 indexed tokenId, uint256 newPrice);

    constructor()
        ERC721("EducationInnovationNFT", "EINFT")
        Ownable(msg.sender)
    {}

    function createInnovation(
        string memory title,
        string memory description,
        uint256 price,
        string memory tokenURI
    ) public returns (uint256) {
        _tokenIds += 1;
        uint256 newItemId = _tokenIds;

        _safeMint(msg.sender, newItemId);
        _setTokenURI(newItemId, tokenURI);

        innovations[newItemId] = Innovation(
            title,
            description,
            price,
            msg.sender,
            true
        );

        creatorInnovations[msg.sender].push(newItemId);

        emit InnovationCreated(newItemId, msg.sender, title);

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

        emit InnovationPurchased(tokenId, buyer, msg.value);
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
}
