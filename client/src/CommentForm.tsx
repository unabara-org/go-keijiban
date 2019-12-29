import React, { useRef } from 'react'
import { CommentDraft } from './index'

type Props = {
  onPost: (draft: CommentDraft) => void
}

export const CommentForm: React.FC<Props> = (props) => {
  const nameInputRef = useRef<HTMLInputElement>(null)
  const textareaRef = useRef<HTMLTextAreaElement>(null)

  const onPost = () => {
    const textAreaElement = textareaRef.current
    const nameInputElement = nameInputRef.current

    if (textAreaElement && nameInputElement) {
      props.onPost({
        nickname: nameInputElement.value,
        body: textAreaElement.value,
      })
    }
  }

  return (
    <div
      style={{
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        backgroundColor: 'rgba(50, 50, 50, 0.4)',
        position: 'absolute',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }}
    >
      <div
        style={{
          backgroundColor: 'white',
          border: '1px solid #eee',
          borderRadius: '4px',
        }}
      >
        <input type={'text'} ref={nameInputRef} placeholder="名前" />
        <textarea ref={textareaRef} placeholder={'コメントを入力してね'} />
        <button onClick={onPost}>投稿する</button>
      </div>
    </div>
  )
}
