import {
    BrowserRouter,
    Routes,
    Route,
} from 'react-router'

import HomePage from '../pages/HomePage/HomePage'
import AreasPage from '../pages/AreasPage/AreasPage'
import ResourcesPage from '../pages/ResoursesPage/ResourcesPage.jsx'

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route
                    path="/"
                    element={<HomePage />}
                />

                <Route
                    path="/areas"
                    element={<AreasPage />}
                />

                <Route
                    path="/resources"
                    element={<ResourcesPage />}
                />
            </Routes>
        </BrowserRouter>
    )
}

export default App