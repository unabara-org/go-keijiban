import React, { useState } from 'react'
import ReactDOM from 'react-dom'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCoffee } from '@fortawesome/free-solid-svg-icons'
import styled from 'styled-components'
import { CommentComponent } from './CommentComponent'
import { CommentForm } from './CommentForm'
import set = Reflect.set

document.addEventListener('DOMContentLoaded', () => {
  const rootEl = document.getElementById('root')

  ReactDOM.render(<App />, rootEl)
})

export type Comment = {
  id: string
  nickname: string
  body: string
  postedAt: Date
}

export type CommentDraft = {
  nickname: string
  body: string
}

function useComments() {
  const [comments, setComments] = useState<Comment[]>([
    {
      id: 'FOo',
      nickname: 'スズキ',
      body: 'テスト投稿',
      postedAt: new Date(),
    },
    {
      id: 'Bar',
      nickname: '88888',
      body: 'BC',
      postedAt: new Date(),
    },
  ])

  async function post(draft: CommentDraft) {
    // ここで API たたく
  }

  return {
    comments,
    setComments,
    post,
  }
}

const App: React.FC = () => {
  const { comments, post } = useComments()
  const [modalIsOpen, setModalIsOpen] = useState<boolean>(false)

  const onPost = (draft: CommentDraft) => {
    post(draft).then(() => {
      setModalIsOpen(false)
    })
  }

  return (
    <div className="flex justify-center h-full bg-gray-100">
      <div className="max-w-lg h-full bg-white relative">
        <Container>
          <HeaderArea>
            <div className="font-bold text-lg px-4">しょぼい掲示板</div>
          </HeaderArea>
          <ContentArea>
            <div
              className="flex flex-col justify-between"
              style={{ marginBottom: '80px' }}
            >
              <div>
                {comments.map((comment) => (
                  <CommentComponent key={comment.id} comment={comment} />
                ))}
              </div>
            </div>
          </ContentArea>
        </Container>
        <div style={{ position: 'absolute', bottom: '16px', right: '16px' }}>
          <div
            className="bg-red-800 rounded-full flex justify-center items-center cursor-pointer"
            style={{ width: '64px', height: '64px' }}
            onClick={() => setModalIsOpen(true)}
          >
            <div>
              <FontAwesomeIcon icon={faCoffee} color="white" />
            </div>
          </div>
        </div>
        {modalIsOpen && <CommentForm onPost={onPost} />}
      </div>
    </div>
  )
}

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
