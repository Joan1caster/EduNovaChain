export type VOTE_TYPE = "Support" | "Opppose" | "Abstain";

export type Table_Basic = {
  index: number;
  name: string;
  publishDate: string;
  sellPrice: string;
  isAi?: boolean;
};

export type TagType = {
  name: string;
  id: number;
};

export type TransactionRecord = {};

export type CitationRecord = {};

export type Comment = {};

export type IdeaInfo = Table_Basic & {
  author: string;
  avator: string;
  topicNameList: string[];
  gradeName: string;
  subjectName: string;
  like: number;
  view: number;
  buy: number;
  userLiked: boolean;
  userCitationed: boolean;
  citation: number;
  rmb: number;
  transactionRecord: TransactionRecord[];
  citationRecord: CitationRecord[];
  comment: Comment[];
};

export type Creator = {
  ID: string;
  Username: string;
  WalletAddress: `0x${string}`;
  CreateAt: string;
};

export type Topic = {
  ID: string;
  Name: string;
  CreateAt: string;
};

export type NFT = {
  Categories: string;
  ContentFeature: string;
  ContractAddress: `0x${string}`;
  CreatedAt: string;
  Creator: Creator;
  CreatorID: number;
  Grades: Topic[];
  ID: number;
  LikeCount: number;
  Owner: Creator;
  OwnerID: number;
  Price: number;
  Subjects: Topic[];
  SummaryFeature: string;
  TokenID: string;
  Topics: Topic[];
  TransactionCount: number;
  UpdatedAt: string;
  ViewCount: number;
  MetadataURI: string; // IPFS Hash
  Title: string;
  Summary: string;
  Content: string;
};
