export interface ImageInterface {
  id: string;
  url: string;
  createdAt: string;
  updatedAt: string;
}

export interface ImageDataInterface{
  images: ImageInterface[] | null;
  currentImage: ImageInterface | null;
  fetchImages: boolean;
}

export interface UserInterface {
  id: string;
  username: string;
  email: string;
  createdAt: string;
  updatedAt: string;
}

export interface GroupListInterface {
  id: string;
  name: string;
  imageUrl: string;
}

export interface GroupInterface {
  id: string;
  name: string;
  imageUrl: string;
  userData: UserInterface[],
  imageData: ImageDataInterface
}

export interface GroupDataInterface {
    groups: GroupListInterface[] | null,
    currentGroup: Partial<GroupInterface> | null,
    fetchGroups: boolean
}