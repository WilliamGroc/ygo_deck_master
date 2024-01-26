import {
  createQuery,
} from '@tanstack/solid-query'
import { http } from './http'
import { Card } from '~/models/card.model'

export const useCards = () => {
  return createQuery(() => ({
    queryKey: ['cards'],
    queryFn: async () => http.get<Card[]>('/cards'),
  }))
}

export const useCard = (id: string) => {
  return createQuery(() => ({
    queryKey: ['cards', id],
    queryFn: async () => http.get<Card>(`/cards/${id}`),
  }))
}