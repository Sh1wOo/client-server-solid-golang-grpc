import { A } from "@solidjs/router";

const Header = () => {
  return (
    <header class="flex w-[95%] mw-[1400px] h-[48px] justify-between items-center m-auto">
      <A href="/">Библиотека</A>
      <nav class="flex gap-2">
        <A href="/все-книги">Все выданые книги</A>
        <A href="/должники">Должники по книгам</A>
      </nav>
    </header>
  );
};

export { Header };
