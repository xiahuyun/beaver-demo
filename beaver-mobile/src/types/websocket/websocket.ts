export interface IWsContent {
  timestamp: number;
  messageId?: string;
  data: {
    type: string;
    body: {
      id: number;
      conversationId: string;
      msg: {
        type: number;
        textMsg: {
          content: string;
        };
        imageMsg: null;
      };
      sender: {
        userId: string;
        avatar: string;
        nickname: string;
      };
      create_at: string;
      msgPreview: string;
    };
  };
}