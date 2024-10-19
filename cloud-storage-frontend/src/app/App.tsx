import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'

import FileSystem from './FileSystem'
import LandingPage from './LandingPage'


const App = () => {
  return (
        <Router>
          <Routes>
            <Route path="/" element={<LandingPage />} />
            <Route path="/filesystem/*" element={<FileSystem />} />
          </Routes>
        </Router>
    )
}

export default App
