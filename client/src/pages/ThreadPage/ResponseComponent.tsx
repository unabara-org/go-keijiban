import React from 'react'
import styled from 'styled-components'
import { Response } from './ThreadPage'

type Props = {
  response: Response
}

export const ResponseComponent: React.FC<Props> = ({ response }) => {
  return (
    <Container>
      <div>
        <div>名前：{response.nickname}</div>
        <div>投稿日時：{response.postedAt.toString()}</div>
      </div>
      <div>{response.body}</div>
    </Container>
  )
}

const Container = styled.div`
  padding: 0.5rem;

  &:not(:last-child) {
    border-bottom: 1px solid #eee;
  }
`
