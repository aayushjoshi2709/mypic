export interface ImageInterface {
  id: string;
  url: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserInterface {
  id: string;
  username: string;
  email: string;
  createdAt: string;
  updatedAt: string;
}


export interface GroupInterface{
  id: string;
  name: string;
  imageUrl: string;
}