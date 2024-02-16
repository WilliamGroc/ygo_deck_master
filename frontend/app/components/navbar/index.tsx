import { Link } from "@remix-run/react";
import { Lang, LangEnum } from "~/const/lang";
import styles from "./styles.module.css";

type Props = {
  setLang: (lang: Lang) => void,
  currentLang: Lang,
  isAuthenticatied: boolean
}

export function Navbar({ setLang, currentLang, isAuthenticatied }: Props) {
  return (
    <div className={styles.navbar}>
      <div className="flex items-center">
        <div className="text-xl font-bold">Yugioh deck master</div>
        <div className="flex flex-row space-x-2 ml-6">
          <Link to="/" className="btn btn-secondary">
            Card list
          </Link>
          {isAuthenticatied ? <Link to="/decks" className="btn btn-secondary">
            Decks
          </Link>
            : <Link to="/login" className="btn btn-secondary">
              Login
            </Link>
          }
        </div>
      </div>
      <div>
        <div>
          <select className="p-2 uppercase" onChange={(e) => setLang(e.target.value as Lang)} value={currentLang}>
            {Object.values(LangEnum).map(lang => (
              <option key={lang} value={lang} className="uppercase">{lang}</option>
            ))}
          </select>
        </div>
      </div>
    </div>
  );
}