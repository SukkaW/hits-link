import type { NextPage } from "next";
import { CopyButton, Leaderboard, Nav, Subtitle, Title } from "../components";
import styled from "styled-components";
import { useState } from "react";

const Container = styled.div`
  width: 80%;
  height: 80vh;
  max-width: 1200px;
  margin: 0 auto;
  @media (max-width: 866px) {
    width: 90%;
  }
`;

const MainContent = styled.div`
  display: flex;
  height: 100%;
  flex-direction: row;

  @media (max-width: 866px) {
    flex-direction: column;
  }
`;

const RightContainer = styled.div`
  display: flex;
  flex-direction: column;
  margin-left: 20px;
  width: 100%;

  @media (max-width: 866px) {
    margin-left: 0;
  }
`;

const RightFlexContainers = styled.div`
  margin: 0;
  padding: 24px 24px;
  flex: 1 1;
  border-radius: 8px;
  background: ${({ theme }) => theme.layoutDark};
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.25);

  &:last-of-type {
    margin-top: 20px;

    @media (max-width: 866px) {
      margin-bottom: 40px;
    }
  }
`;

const Input = styled.input`
  width: 100%;
  min-width: 75%;
  max-width: 400px;
  margin: 10px 0 25px 0;
  padding: 15px 10px;
  font-size: 1em;
  border: none;
  border-radius: 8px;
  color: ${({ theme }) => theme.textLight};
  background: ${({ theme }) => theme.layoutLittleLessDark};
  outline: 0;
  box-shadow: 0px 0px 12px 0px rgb(0 0 0 / 16%);
  transition: all 0.2s ease-in-out;

  &:focus {
    box-shadow: 0px 0px 14px rgba(0, 0, 0, 0.25);
  }

  &:focus::placeholder {
    color: ${({ theme }) => theme.textLight}9d;
  }
`;

const InputTitle = styled.h3`
  margin: 0;
  padding: 0;
  font-size: 1em;
  font-weight: 700;
  color: ${({ theme }) => theme.textBlackest};
`;

const InputContainer = styled.div`
  position: relative;
  display: inline-flex;
  width: 100%;
  max-width: 80%;
`;

const Home: NextPage = () => {
  const [URL, setURL] = useState("");

  return (
    <Container>
      <Nav />
      <MainContent>
        <Leaderboard />
        <RightContainer>
          <RightFlexContainers>
            <Title>Create a Hit</Title>
            <Subtitle>Generate a SVG Image for your link</Subtitle>
            <br />
            <Input
              style={{ maxWidth: 300 }}
              placeholder="Enter your URL"
              onChange={(e) => setURL(e.target.value)}
            />
          </RightFlexContainers>
          <RightFlexContainers>
            <Title>Copy URL</Title>
            <Subtitle>Choose from the following:</Subtitle>
            <br />

            <InputTitle>HTML</InputTitle>
            <InputContainer>
              <Input
                style={{
                  padding: "15px 87px 15px 10px",
                }}
                value={`<img src="${
                  process.env.NODE_ENV === "development"
                    ? "localhost:3000"
                    : "https://hits.link"
                }/hits?url=${URL}" />`}
                disabled={true}
              />
              <CopyButton
                text={`<img src="${
                  process.env.NODE_ENV === "development"
                    ? "localhost:3000"
                    : "https://hits.link"
                }/hits?url=${URL}" />`}
              />
            </InputContainer>

            <InputTitle>Markdown</InputTitle>
            <InputContainer>
              <Input
                style={{
                  padding: "15px 87px 15px 10px",
                }}
                value={`![Hits](${
                  process.env.NODE_ENV === "development"
                    ? "localhost:3000"
                    : "https://hits.link"
                }/hits?url=${URL})`}
                disabled={true}
              />
              <CopyButton
                text={`![Hits](${
                  process.env.NODE_ENV === "development"
                    ? "localhost:3000"
                    : "https://hits.link"
                }/hits?url=${URL})`}
              />
            </InputContainer>
          </RightFlexContainers>
        </RightContainer>
      </MainContent>
    </Container>
  );
};

export default Home;
