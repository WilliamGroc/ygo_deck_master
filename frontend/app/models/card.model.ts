export type CardLang = {
  name: string;
  effectText: string;
}

export type Card = {
  id: number;
  name: string;
  type: string;
  frameType: string;
  race: string;
  level: number;
  attribute: string;
  linkVal: number;
  atk: number;
  def: number;
  fr: CardLang;
  en: CardLang;
  de: CardLang;
  it: CardLang;
  es: CardLang;
  ja: CardLang;
  ko: CardLang;
  pt: CardLang;
};
