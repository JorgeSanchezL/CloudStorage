import React, { createContext, ReactNode, useContext, useState } from "react";
import { createPortal } from "react-dom";

import useCreatePortal from "./useCreatePortal";

interface ModalContextType {
    isModalOpen: boolean;
    openModal: (content: ReactNode) => void;
    closeModal: () => void;
}

export const ModalContext = createContext<ModalContextType | null>(null);

export const useModal = (): ModalContextType => {
    const context = useContext(ModalContext);
    if (!context) {
        throw new Error('useModal must be used within a ModalProvider');
    }
    return context;
};

export const ModalProvider: React.FC<{children: React.ReactNode}> = ({ children }) => {
    const [isModalOpen, setIsModalOpen] = useState(false);
    
    const {portalRoot} = useCreatePortal();
    const [modalContent, setModalContent] = useState(null);

    const openModal = (componentToRender) => {
        setIsModalOpen(true);
        setModalContent(componentToRender);
    }
    const closeModal = () => {
        setIsModalOpen(false);
    }
    

    return (
        <ModalContext.Provider value={{isModalOpen, openModal, closeModal}}>
            {children}
            {isModalOpen && modalContent && createPortal(
                modalContent,
                portalRoot as Element,
            )}
        </ModalContext.Provider>
    );
}