import './style.css'
import { CopyToClipboard } from '@wailsjs/go/main/App'

function HomePage({
  textoClipboard = ''
}) {

  const handleClick = async () => {
    const content = document.querySelector('.bodyContent').textContent
    const resultado = await CopyToClipboard(content)
    console.log(resultado)
  }

  return (
    < div className='content' >
      <div className='header'>
        <h1>Clipboard Manager</h1>
        <span>Gerencie textos e inputs localmente</span>
      </div>

      <div className='body'>
        <div className='bodyContent'>
          {textoClipboard || 'Aguardando informações...'}
        </div>
      </div>

      <div className='footer'>
        <button onClick={handleClick}>Copiar Conteúdo</button>
      </div>
    </div>
  )
}

export default HomePage
