import { useState } from 'react'
import './App.css'

function App() {

  const [responseText, setResponseText] = useState('')

  async function getText() {
    try {
      const response = await fetch('/api')

      if (!response.ok) {
        throw new Error(`HTTP error: ${response.status}`)
      }

      const text = await response.text()
      setResponseText(text)
    } catch (error) {
      console.error(error)
      setResponseText("Failed requesting to server.")
    }
  }

  return (
      <main className="container">
        <button type="button"
                onClick={getText}>Get text</button>

        <textarea value={responseText}
                  readOnly
                  placeholder="Nothing"/>
      </main>
  )

}

export default App