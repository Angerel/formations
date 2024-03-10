export class FetchService {
    static fetch<T>(url:string) : Promise<T> {
        return fetch(url)
            .then((response)=>{
                    return <T>response.body;
                }
            );
    }
}