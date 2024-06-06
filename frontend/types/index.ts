export interface Invitation {
    id: number
    email: string
    status: string
  }
  
  export interface Registration {
    email: string
    nama: string
    tanggalLahir: string
    jenisKelamin: string
    alamatLengkap: string
    registrationCode: string
    qrCode: string
  }
  