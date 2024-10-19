import { useEffect } from "react";

const useCreatePortal = (modalId = 'portal', tagName = 'div') => {
    const portalRoot = document.getElementById(modalId);
    const portalElement = document.createElement(tagName);

    useEffect(() => {
        if (portalRoot === null) {
            portalElement.setAttribute('id', modalId);
            document.body.appendChild(portalElement);
        }

        return () => {
            if (portalRoot !== null) {
                document.body.removeChild(portalRoot);
            }
        }
    }, [modalId, tagName])

    return {modalId, portalRoot}
}

export default useCreatePortal