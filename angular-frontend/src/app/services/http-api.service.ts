import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { environment } from '@env';
import { catchError, Observable, throwError } from 'rxjs';
import { Router } from '@angular/router';

@Injectable({ providedIn: 'root' })
export class HttpAPI {
  private readonly API_URL = environment.apiUrl;

  constructor(
    private http: HttpClient,
    private router: Router,
  ) {}

  private getHeaders(): HttpHeaders {
    const token = localStorage.getItem('token');

    let headers = new HttpHeaders({
      'Content-Type': 'application/json',
    });

    if (token) {
      headers = headers.set('Authorization', `Bearer ${token}`);
    }
    return headers;
  }

  private handleError(error: any) {
    if (error.status === 401) {
      localStorage.clear(); 
      this.router.navigate(['/login']);
    }
    return throwError(() => error);
  }

  get<T>(api: string, params?: HttpParams): Observable<T> {
    return this.http
      .get<T>(`${this.API_URL}${api}`, { headers: this.getHeaders(), params })
      .pipe(catchError((err) => this.handleError(err)));
  }

  post<T>(api: string, body: any): Observable<T> {
    return this.http
      .post<T>(`${this.API_URL}${api}`, body, { headers: this.getHeaders() })
      .pipe(catchError((err) => this.handleError(err)));
  }

  put<T>(api: string, body: any, id?: string | number): Observable<T> {
    const url = id ? `${this.API_URL}${api}/${id}` : `${this.API_URL}${api}`;
    return this.http
      .put<T>(url, body, { headers: this.getHeaders() })
      .pipe(catchError((err) => this.handleError(err)));
  }

  patch<T>(api: string, body: any): Observable<T> {
    return this.http
      .patch<T>(`${this.API_URL}${api}`, body, { headers: this.getHeaders() })
      .pipe(catchError((err) => this.handleError(err)));
  }

  delete<T>(api: string, id?: string | number): Observable<T> {
    const url = id ? `${this.API_URL}${api}/${id}` : `${this.API_URL}${api}`;
    return this.http
      .delete<T>(url, { headers: this.getHeaders() })
      .pipe(catchError((err) => this.handleError(err)));
  }
  deleteWithBody<T>(api: string, body: any): Observable<T> {
    return this.http
      .delete<T>(`${this.API_URL}${api}`, {
        headers: this.getHeaders(),
        body: body,
      })
      .pipe(catchError((err) => this.handleError(err)));
  }
}
