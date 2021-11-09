import service from '@/utils/request-upload'

 
export const uploadFile = (data) => {
  return service({   
    url: '/commFile/upload',
    method: 'post',    
    data
  })
}
 