import { createStore } from "solid-js/store";
import useBooks from "../../hooks/useBooks";
import { useNavigate } from "@solidjs/router";

interface BookData {
  name: string;
  author: string;
  year: string;
}

const AddBook = () => {
  const [data, setData] = createStore<BookData>({
    name: "",
    author: "",
    year: "",
  });
  const { refetchBooks } = useBooks();
  const navigate = useNavigate()

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement>,
    field: keyof BookData
  ) => {
    setData({ ...data, [field]: e.target.value });
  };

  const handleOnSubmit = async (e: Event) => {
    e.preventDefault();

    await fetch("http://localhost:8090/books/add", {
      method: "POST",
      body: JSON.stringify({
        title: data.name,
        author: data.author,
        year: +data.year,
      }),
    }).then(refetchBooks);

    // Очистка данных после отправки
    setData({
      name: "",
      author: "",
      year: "",
    });

    navigate('/все-книги')
  };

  return (
    <form
      onSubmit={handleOnSubmit}
      class="w-[350px] flex flex-col border p-4 m-auto rounded text-center"
    >
      <h2 class="pb-2 text-xl">Добавление книги</h2>
      <input
        type="text"
        value={data.name}
        onChange={(e) => handleInputChange(e, "name")}
        placeholder="Название книги"
        class="border p-2 mb-2 rounded"
      />
      <input
        type="text"
        value={data.author}
        onChange={(e) => handleInputChange(e, "author")}
        placeholder="Автор книги"
        class="border p-2 mb-2 rounded"
      />
      <input
        type="number"
        value={data.year}
        onChange={(e) => handleInputChange(e, "year")}
        placeholder="Год книги"
        class="border p-2 mb-2 rounded"
      />
      <button
        type="submit"
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      >
        Добавить книгу
      </button>
    </form>
  );
};

export default AddBook;
