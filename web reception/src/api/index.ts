import request from '../utils/request';

const HOST_SERVER = 'http://192.168.50.22:1232'


export const setgo = ({headers}: { headers: any }) => {
    return request({
        url: HOST_SERVER + '/api/v2/setgo',
        method: 'get',
        headers
    });
}


