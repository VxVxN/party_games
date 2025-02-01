import {QueryClient, QueryClientProvider} from '@tanstack/react-query'
import './styles/App.css'
import {BrowserRouter, Route, Routes} from "react-router";
import HomePage from "./pages/HomePage.tsx";
import GamePage from "./pages/GamePage.tsx";
import MainLayout from "./layouts/MainLayout.tsx";

const queryClient = new QueryClient()

function App() {

    return (
        <QueryClientProvider client={queryClient}>
            <BrowserRouter>
                <Routes>
                    <Route element={<MainLayout/>}>
                        <Route index path="/" element={<HomePage/>}/>
                        <Route path="/game/:topic" element={<GamePage/>}/>
                    </Route>
                </Routes>
            </BrowserRouter>
        </QueryClientProvider>
    )
}

export default App
