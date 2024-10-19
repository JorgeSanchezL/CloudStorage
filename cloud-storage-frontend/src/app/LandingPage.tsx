import React from 'react';
import { useNavigate } from 'react-router-dom';
import './LandingPage.css';

const LandingPage: React.FC = () => {
    const navigate = useNavigate();

    return (
        <div className="container">
            <div className="shape shape1"></div>
            <div className="shape shape2"></div>
            <div className="shape shape3"></div>

            <div className="content-box">
                <h1 className="heading">CloudStorage</h1>
                <p className="paragraph">Store and manage your files securely in a private cloud.</p>
                <button 
                    onClick={() => navigate('/filesystem/')} 
                    className="button"
                >
                    Go to FileSystem
                </button>
            </div>
        </div>
    );
};

export default LandingPage;