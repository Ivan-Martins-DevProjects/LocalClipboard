import { useEffect, useState } from 'react';
import { EventsOn } from '@wailsjs/runtime/runtime';
import './App.css';
import HomePage from './pages/home';

function App() {
  const [texto, setTexto] = useState("")

  useEffect(() => {
    const setypWs = () => {
      console.log('Iniciando conexão com servidor WebSocket...')

      const quit = EventsOn("novo_clipboard", (message) => {
        setTexto(message)
      })

      return quit
    }

    const stopListening = setypWs()

    return () => stopListening()
  }, [])

  return (
    <div className='main'>
      <HomePage textoClipboard={texto} />
    </div>
  )
}

export default App
