import React from 'react'
import styled from 'styled-components'
import { Comment } from './index'

type Props = {
  comment: Comment
}

export const CommentComponent: React.FC<Props> = ({ comment }) => {
  return (
    <Container>
      <div>
        <div>名前：{comment.nickname}</div>
        <div>投稿日時：{comment.postedAt.toString()}</div>
      </div>
      <div>{comment.body}</div>
    </Container>
  )
}

const Container = styled.div`
  padding: 16px;
  cursor: pointer;

  &:not(:last-child) {
    border-bottom: 1px solid #eee;
  }
`
