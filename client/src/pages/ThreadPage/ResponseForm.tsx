import React, { useRef } from 'react'
import { ResponseDraft } from './ThreadPage'

type Props = {
  onPost: (draft: ResponseDraft) => void
  onRequestClose: () => void
}

export const ResponseForm: React.FC<Props> = (props) => {
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

  function onClickOutside(e: React.MouseEvent<HTMLDivElement>) {
    props.onRequestClose()
  }

  function onClickInside(e: React.MouseEvent<HTMLDivElement>) {
    e.stopPropagation()
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
      onClick={onClickOutside}
    >
      <div
        style={{
          backgroundColor: 'white',
          border: '1px solid #eee',
          borderRadius: '4px',
        }}
        onClick={onClickInside}
      >
        <input type={'text'} ref={nameInputRef} placeholder="名前" />
        <textarea ref={textareaRef} placeholder={'コメントを入力してね'} />
        <button onClick={onPost}>投稿する</button>
      </div>
    </div>
  )
}
