import './App.css'


import NavBar from '../components/NavBar'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import FileSystem from './FileSystem'


const App = () => {
  return (
        <Router>
          <NavBar text='CloudStorage' />
          <Routes>
            <Route path="/:path/*" element={<FileSystem setParentLoad={(_: boolean) => {}} parentPath=''/>} />
          </Routes>
        </Router>
    )
}

export default App
