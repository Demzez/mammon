import { useState } from 'react'
import '../pages.css'
import {Link} from "react-router";

function ResourcesPage({ onBack }) {
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

            <Link
                to="/"
                className="button-link"
            >
                Back
            </Link>
        </main>
    )
}

export default ResourcesPage