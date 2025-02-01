import {motion} from "framer-motion";
import './styles.scss'
import React from "react";
import {Link} from "react-router";

interface GameListItemProps {
    title?: string
    topic: string
    children: React.ReactNode
}

function GameListItem(props: GameListItemProps) {
    const {title, children, topic} = props;

    return <Link to={`/game/${topic}`}>
        <motion.li className='game-item' initial={{x: '100%', scale: 1}}
                   animate={{x: 0}}
                   whileHover={{scale: 1.05, backgroundColor: 'var(--secondary)'}}
                   whileTap={{scale: 0.95}}>
            {title || children}
        </motion.li>
    </Link>
}

export default GameListItem;