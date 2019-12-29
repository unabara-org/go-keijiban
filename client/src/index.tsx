import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter, Switch, Route, Link } from 'react-router-dom'
import { ThreadListPage } from './pages/ThreadListPage/ThreadListPage'
import styled from 'styled-components'
import { ThreadPage } from './pages/ThreadPage/ThreadPage'

document.addEventListener('DOMContentLoaded', () => {
  const rootEl = document.getElementById('root')

  ReactDOM.render(<App />, rootEl)
})

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <div className="flex justify-center h-full bg-gray-100">
        <Wrapper className="h-full bg-white relative">
          <Container>
            <HeaderArea>
              <div className="px-4">
                <Link to="/">
                  <Title>しょぼい掲示板</Title>
                </Link>
              </div>
            </HeaderArea>
            <ContentArea>
              <div
                className="flex flex-col justify-between"
                style={{ marginBottom: '80px' }}
              >
                <Switch>
                  <Route exact path="/">
                    <ThreadListPage />
                  </Route>
                  <Route path="/threads/:id">
                    <ThreadPage />
                  </Route>
                </Switch>
              </div>
            </ContentArea>
          </Container>
        </Wrapper>
      </div>
    </BrowserRouter>
  )
}

const Wrapper = styled.div`
  width: 640px;

  @media (max-width: 640px) {
    width: 100%;
  }
`

const Container = styled.div`
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: 48px 1fr;
  grid-template-areas: 'Header' 'Content';

  position: relative;
  height: 100%;
`

const HeaderArea = styled.div`
  grid-area: Header;
`

const ContentArea = styled.div`
  grid-area: Content;
  overflow: scroll;
`

const Title = styled.span`
  font-weight: bold;

  // これを参考にした
  // https://spyweb.media/2017/11/16/css-hover-text-highlighter-line-animation/
  background-image: linear-gradient(transparent 60%, yellow 40%);
  background-position: 200% 0%;
  background-size: 200% auto;
  background-repeat: no-repeat;

  transition: background-position 0.3s ease-out;

  &:hover {
    background-position: 0% 0%;
  }
`
