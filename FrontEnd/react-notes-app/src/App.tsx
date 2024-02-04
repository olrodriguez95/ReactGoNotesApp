import { Container, Row } from "react-bootstrap";
import Navbar from "./components/Navbar";
import NoteCard from "./components/NoteCard";
import { useEffect, useState } from "react";

interface Note {
  noteTitle: String;
  noteText: String;
  starred: boolean;
}

function App() {
  const newNote: Note = { noteTitle: "", noteText: "", starred: false };
  const [note, setNote] = useState<Note>(newNote);

  const mapNote = (data: any) => {
    var tempNote: Note = {
      noteText: data[0].note_text,
      noteTitle: data[0].note_title,
      starred: data[0].favorite,
    };

    setNote(tempNote);
  };

  useEffect(() => {
    const getNotes = () => {
      fetch("http://127.0.0.1:8080/notes", {
        method: "GET",
      })
        .then((response) => response.json())
        .then((data) => {
          mapNote(data);
          console.log(data);
        })
        .catch((error) => console.log(error));
    };
    console.log("Running Use Effect");
    getNotes();
  }, []);

  return (
    <>
      <Navbar />
      <Container>
        <Row>
          {
            <NoteCard
              noteTitle={note?.noteTitle ?? "NO TITLE"}
              noteText={note?.noteText ?? "NO TEXT"}
              starred={note?.starred ?? false}
            />
          }
        </Row>
      </Container>
    </>
  );
}

export default App;
