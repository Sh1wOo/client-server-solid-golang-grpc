import { createEffect } from "solid-js";
import { createStore } from "solid-js/store";

export type TGetBooks = Array<{
  id: string;
  title: string;
  author: string;
  year: number;
}>;

export default function useBooks() {
  const [books, setBooks] = createStore<TGetBooks>([]);

  const getBooks = async () => {
    const res = await fetch("http://localhost:8090/books");
    const data: TGetBooks = await res.json();
    
    // console.log(data);
    setBooks(data);
  };

  const refetchBooks = async () => {
    await getBooks();
  };

  createEffect(() => {
    getBooks();
  }, []);

  return { books, refetchBooks };
}
