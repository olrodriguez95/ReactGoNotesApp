import { useEffect, useState } from "react";
import { BsStar, BsStarFill } from "react-icons/bs";

interface StarToggleProps {
  starred: boolean;
}

function StarToggleButton({ starred }: StarToggleProps) {
  const [isStarred, setIsStarred] = useState(starred);

  const handleToggle = () => {
    setIsStarred(!isStarred);
  };

  return (
    <button
      className={`btn ${isStarred ? "btn-warning" : "btn-light"}`}
      onClick={handleToggle}
    >
      {isStarred ? <BsStarFill /> : <BsStar />}
    </button>
  );
}

export default StarToggleButton;
