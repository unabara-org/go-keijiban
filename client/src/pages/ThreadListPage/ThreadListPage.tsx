import React, { useState, useEffect } from 'react'
import styled from 'styled-components'
import { Link } from 'react-router-dom'

type Props = {}

type ThreadMeta = {
  id: string
  title: string
  responseCount: number
  createAt: Date
}

export const ThreadListPage: React.FC<Props> = () => {
  const { threadMetas, isLoading } = useThreadMeta()

  return (
    <div>
      {threadMetas.map((meta) => (
        <Container to={`/threads/${meta.id}`} key={meta.id}>
          <div>{meta.createAt.toDateString()}</div>
          <div className="flex justify-between">
            <div className="font-bold text-lg">{meta.title}</div>
            <div
              style={{
                width: '3rem',
                marginLeft: '1rem',
                textAlign: 'right',
                alignSelf: 'center',
              }}
            >
              {meta.responseCount}
            </div>
          </div>
        </Container>
      ))}
    </div>
  )
}

function useThreadMeta() {
  const [threadMetas, setThreadMetas] = useState<ThreadMeta[]>([])
  const [isLoading, setIsLoading] = useState<boolean>(false)

  useEffect(() => {
    setThreadMetas([
      {
        id: '99',
        title: '今年ガチで買って良かったものあげてけ',
        responseCount: 780,
        createAt: new Date(),
      },
      {
        id: '88',
        title:
          '格安SIMの会社が乱立してて惑乱してる　乗換えたいんだが、とりあえずUQモバイルでいいのか？情つよモメンたろすけ',
        responseCount: 343,
        createAt: new Date(),
      },
    ])
  }, [])

  return {
    threadMetas,
    isLoading,
  }
}

const Container = styled(Link)`
  cursor: pointer;
  display: block;
  padding: 0.5rem;

  &:hover {
    background-color: #f5f5f5;
  }

  &:not(:last-child) {
    border-bottom: 1px solid #eee;
  }
`
