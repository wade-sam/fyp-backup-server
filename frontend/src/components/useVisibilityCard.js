import React, { useState } from 'react'

const useVisibilityCard = (component, visibility = false) => {
    const [visible, setVisibility] = useState(() => visibility);
    return [visible ? component : null, () => setVisibility((v) => !v)];
}
export default useVisibilityCard;