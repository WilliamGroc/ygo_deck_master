import { Card } from "./card.model";

export interface Deck {
  id: number;
  name: string;
  cards: Card[];
  createdBy: number;
  createdAt: string;
  updatedAt: string;
  isPublic: boolean;
}