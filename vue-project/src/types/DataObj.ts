
export interface PArray {
    p: number;
}

export interface Options {
    id: number;
    text: string;
}

export interface DataObj {
    data: PArray[];
    type: string;
    options: Options[];
}
