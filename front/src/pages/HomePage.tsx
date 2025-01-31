import PageHeader from "../components/PageHeader";
import GameList from "../components/GameListCategory/GameList.tsx";
import {motion} from "framer-motion";
import VideoBackground from "../components/VideoBackground";

function HomePage() {

    return <div className='page'>
        <div className='relative top-0 left-0 w-full h-full z-9 rounded-lg p-4 flex flex-col justify-center content-center items-center gap-2'
             style={{backgroundColor: 'rgba(15,19,32,0.8)'}}>
            <PageHeader title='Party games'/>
            <motion.div initial={{opacity: .5}} animate={{opacity: 1}}
                        transition={{duration: 1, repeat: Infinity, repeatType: 'mirror'}}>Please
                select the category game
            </motion.div>
            <GameList/>
        </div>
        <VideoBackground/>
    </div>
}

export default HomePage;