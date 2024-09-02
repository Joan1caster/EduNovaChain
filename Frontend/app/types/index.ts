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
  key: number;
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
