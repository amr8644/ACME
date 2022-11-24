import React, { useState } from "react";
import { Modal, Button, Text, Textarea } from "@nextui-org/react";
import { API_URL } from "../pages/api/url";
import Pen from "./icons/Pen";
import { addPost } from "../pages/api/post";
import { useMutation, useQueryClient } from "react-query";

export default function AddPostModal() {
  const handler = () => setVisible(true);
  const [visible, setVisible] = React.useState(false);
  const queryClient = useQueryClient();

  let user;
  if (typeof window !== "undefined") {
    user = JSON.parse(localStorage.getItem("user") ?? "");
  } else {
    console.log("You are on the server");
  }
  const userid = user?.ID;

  const [content, setContent] = useState("");

  const closeHandler = () => {
    setVisible(false);
  };

  const addPostMutation = useMutation(addPost, {
    onSuccess: () => {
      queryClient.invalidateQueries("posts");
    },
  });

  return (
    <>
      <form action={`${API_URL}/register`} method="POST" id="postform">
        <div>
          <style jsx>{`
            div {
              position: relative;
            }
          `}</style>
          <Button
            css={{
              position: "fixed",
              bottom: 100,
              right: 160,
              borderRadius: 9999,
              width: "60px",
              height: "60px",
            }}
            auto
            flat
            color={"primary"}
            onClick={handler}
          >
            <Pen fill="currentColor" />
          </Button>
          <Modal
            blur
            closeButton
            width="500px"
            aria-labelledby="modal-title"
            open={visible}
            onClose={closeHandler}
          >
            <Modal.Header>
              <Text id="modal-title" size={18}>
                Write
              </Text>
            </Modal.Header>

            <Modal.Body>
              <Textarea
                bordered
                name="Content"
                fullWidth
                color="primary"
                size="lg"
                placeholder="Content"
                form="postform"
                onChange={(e) => setContent(e.target.value)}
              />
            </Modal.Body>
            <Modal.Footer>
              <Button auto flat color="error" onClick={closeHandler}>
                Close
              </Button>
              <Button
                onClick={() => {
                  closeHandler();
                  addPostMutation.mutate({ id: userid, content: content });
                }}
                flat
                type="submit"
                auto
                color={"success"}
              >
                Post
              </Button>
            </Modal.Footer>
          </Modal>
        </div>
      </form>
    </>
  );
}
