export interface BookV1 {
  id: number;
  title: string;
  author: string;
  owner_id: number;
}

export interface BookV2 {
  _id: string;
  title: string;
  author: string;
  owner_id: number;
}

export interface User {
  id: number;
  name: string;
  email: string;
}
