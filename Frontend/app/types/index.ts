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
