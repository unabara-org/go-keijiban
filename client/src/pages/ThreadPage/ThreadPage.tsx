import React, { useState } from 'react'
import { useParams } from 'react-router-dom'
import { ResponseComponent } from './ResponseComponent'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlus } from '@fortawesome/free-solid-svg-icons'
import { ResponseForm } from './ResponseForm'

type PathParams = {
  id: string
}

export const ThreadPage: React.FC = () => {
  const { id } = useParams<PathParams>()
  const { threadWithResponses, post } = useThreadWithResponses(id)
  const [modalIsOpen, setModalIsOpen] = useState<boolean>(false)

  return (
    <div>
      <div
        className="flex flex-col justify-between"
        style={{ marginBottom: '80px' }}
      >
        <div>
          {threadWithResponses.responses.map((response) => (
            <ResponseComponent key={response.id} response={response} />
          ))}
        </div>
      </div>
      <div style={{ position: 'absolute', bottom: '16px', right: '16px' }}>
        <div
          className="bg-red-800 rounded-full flex justify-center items-center cursor-pointer"
          style={{ width: '64px', height: '64px' }}
          onClick={() => setModalIsOpen(true)}
        >
          <div>
            <FontAwesomeIcon icon={faPlus} color="white" />
          </div>
        </div>
      </div>
      {modalIsOpen && (
        <ResponseForm
          onPost={post}
          onRequestClose={() => setModalIsOpen(false)}
        />
      )}
    </div>
  )
}

type ThreadWithResponses = {
  id: string
  title: string
  createdAt: Date
  responses: Response[]
}

export type Response = {
  id: string
  nickname: string
  body: string
  postedAt: Date
}

export type ResponseDraft = {
  nickname: string
  body: string
}

function useThreadWithResponses(threadId: string) {
  const [threadWithResponses] = useState<ThreadWithResponses>({
    id: '99',
    title: '今年ガチで買って良かったものあげてけ',
    createdAt: new Date(),
    responses: [
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
    ],
  })

  async function post(draft: ResponseDraft) {
    // ここで API たたく
  }

  return {
    threadWithResponses,
    post,
  }
}
