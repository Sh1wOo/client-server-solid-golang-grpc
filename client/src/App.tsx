import { Header } from './components/header'
import './App.css'

function App(props: any) {

  return (
    <>
      <Header/>
      <main class='flex flex-col w-[95%] mw-[1400px] h-[calc(100vh_-48px)] justify-between items-center m-auto'>
        {props.children}
      </main>
    </>
  )
}

export default App
