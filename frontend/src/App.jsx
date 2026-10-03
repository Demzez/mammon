import { useState } from 'react'
import './App.css'

function App() {
  const [page, setPage] = useState('home')
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
      setResponseText('Failed requesting to server.')
    }
  }

  if (page === 'areas') {
    return (
        <main className="container">
          <button
              type="button"
              onClick={getText}
          >
            Get text
          </button>

          <textarea
              value={responseText}
              readOnly
              placeholder="Nothing"
          />

          <button
              type="button"
              onClick={() => setPage('home')}
          >
            Back
          </button>
        </main>
    )
  }

  return (
      <main className="menu-page">
        <div className="menu-grid">
          <button
              className="menu-card"
              onClick={() => setPage('areas')}
          >
            Areas
          </button>

          <button className="menu-card">
            Projects
          </button>

          <button className="menu-card">
            Archive
          </button>

          <button className="menu-card">
            Resources
          </button>
        </div>
      </main>
  )
}

export default App