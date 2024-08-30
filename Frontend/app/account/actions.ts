'use server'
import { redirect } from "next/navigation"
import { cookies } from "next/headers"

export async function login(prevState, formData) {

    const data = {
        email: formData.get('email'),
        password: formData.get('password')
    }

    console.log(data)
    cookies().set("email", data.email)
    redirect('/')
    return {}
}

export async function signup(prevState, formData) {

    const data = {
        email: formData.get('email'),
        password: formData.get('password')
    }

    console.log(data)
    cookies().set("email", data.email)
    redirect('/')
    return {}
}