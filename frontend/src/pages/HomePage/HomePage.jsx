import { Link } from 'react-router'

import '../pages.css'

function HomePage() {
    return (
        <main className="menu-page">
            <div className="menu-grid">

                <Link
                    className="menu-card"
                    to="/areas"
                >
                    Areas
                </Link>

                <button className="menu-card">
                    Projects
                </button>

                <button className="menu-card">
                    Archive
                </button>

                <Link
                    className="menu-card"
                    to="/resources"
                >
                    Resources
                </Link>

            </div>
        </main>
    )
}

export default HomePage