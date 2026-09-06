import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { AuthService } from '@services/auth.service';
import { SqlCleanDirective } from '../../directives/sql-clean.directive';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, RouterLink, SqlCleanDirective],
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  private fb = inject(FormBuilder);
  private authService = inject(AuthService);
  private router = inject(Router);

  registerForm!: FormGroup;
  showPassword = false;
  showConfirmPassword = false;
  isSubmitted = false;
  passwordMismatch = false;

  ngOnInit(): void {
    this.registerForm = this.fb.group({
      username: ['', [Validators.required, Validators.minLength(3), Validators.pattern('^[a-zA-Z0-9_-]+$')]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      confirmPassword: ['', [Validators.required]]
    });
  }

  togglePassword(): void {
    this.showPassword = !this.showPassword;
  }

  toggleConfirmPassword(): void {
    this.showConfirmPassword = !this.showConfirmPassword;
  }

  isFieldInvalid(fieldName: string): boolean {
    const field = this.registerForm.get(fieldName);

    if (fieldName === 'password' || fieldName === 'confirmPassword') {
      return !!((field && field.invalid && (field.touched || this.isSubmitted)) || this.passwordMismatch);
    }

    return !!(field && field.invalid && (field.touched || this.isSubmitted));
  }

  onSubmit(): void {
    this.isSubmitted = true;
    this.passwordMismatch = false;

    if (this.registerForm.invalid) {
      Swal.fire({
        icon: 'warning',
        title: 'ข้อมูลไม่ถูกต้อง',
        text: 'กรุณากรอกข้อมูลให้ครบถ้วนและใช้อักขระที่อนุญาตเท่านั้น',
        confirmButtonText: 'ตกลง',
        confirmButtonColor: '#4CAF50'
      });
      return;
    }

    if (this.registerForm.value.password !== this.registerForm.value.confirmPassword) {
      this.passwordMismatch = true;
      Swal.fire({
        icon: 'error',
        title: 'รหัสผ่านไม่ตรงกัน',
        text: 'กรุณาตรวจสอบรหัสผ่านใหม่อีกครั้ง',
        confirmButtonText: 'ตกลง',
        confirmButtonColor: '#e53935'
      });
      return;
    }

    this.authService.register(this.registerForm.value).subscribe({
      next: () => {
        Swal.fire({
          icon: 'success',
          title: 'สมัครสมาชิกสำเร็จ',
          text: 'ระบบกำลังนำท่านไปยังหน้าเข้าสู่ระบบ',
          timer: 1500,
          showConfirmButton: false
        }).then(() => {
          this.router.navigate(['/login']);
        });
      },
      error: () => {
        Swal.fire({
          icon: 'error',
          title: 'สมัครสมาชิกไม่สำเร็จ',
          text: 'ชื่อผู้ใช้งานนี้มีอยู่ในระบบแล้ว',
          confirmButtonText: 'ตกลง',
          confirmButtonColor: '#e53935'
        });
      }
    });
  }
}
