import axios from "axios";
import { GetServerSideProps } from "next";
import AddPostModal from "../components/AddPostModal";
import PostCard from "../components/Card";
import { API_URL } from "./api/url";
import NavBar from "../components/NavBar";
import News from "../components/News";
import { Container } from "@nextui-org/react";

export default function Dashboard({ posts }: any) {
   return (
      <>
         <NavBar />
         <AddPostModal />
         <PostCard posts={posts} />
      </>
   );
}

export const getServerSideProps: GetServerSideProps = async (context: any) => {
   const res = await axios.get(`${API_URL}/dashboard`, {
      withCredentials: true,
   });
   const posts = res.data;

   return {
      props: {
         posts: JSON.parse(JSON.stringify(posts)),
      },
   };
};
