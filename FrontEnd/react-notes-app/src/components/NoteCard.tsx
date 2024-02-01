import StarToggleButton from "./StarToggleButton";

interface NoteProps {
  noteTitle?: String;
  noteText: String;
  starred: boolean;
}

function NoteCard({ noteTitle, noteText, starred }: NoteProps) {
  return (
    <>
      <div className="card" style={{ width: "18rem", margin: "1rem" }}>
        <StarToggleButton starred={starred} />
        <div className="card-body">
          <h5 className="card-title">{noteTitle}</h5>
          <p className="card-text">{noteText}</p>
          <a href="#" className="btn btn-primary">
            View Note
          </a>
        </div>
      </div>
    </>
  );
}

export default NoteCard;
