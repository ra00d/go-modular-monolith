import {
  createContext,
  useState,
  useMemo,
  useContext,
  type ReactNode,
  type Dispatch,
  type SetStateAction,
} from "react";

export interface BreadcrumbsValue {
  parts: string[];
  current: string;
  setCurrent: Dispatch<SetStateAction<string>>;
  setParts: Dispatch<SetStateAction<string[]>>;
}

export const BreadcrumbsContext = createContext<BreadcrumbsValue | null>(null);

export function BreadcrumbsProvider({ children }: { children: ReactNode }) {
  const [parts, setParts] = useState<string[]>([]);
  const [current, setCurrent] = useState<string>("");
  return (
    <BreadcrumbsContext.Provider
      value={{
        parts,
        setParts,
        current,
        setCurrent,
      }}
    >
      {children}
    </BreadcrumbsContext.Provider>
  );
}
