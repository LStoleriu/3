package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for reducemaxvecnorm2 kernel
var reducemaxvecnorm2_code cu.Function

// Stores the arguments for reducemaxvecnorm2 kernel invocation
type reducemaxvecnorm2_args_t struct {
	arg_x       unsafe.Pointer
	arg_y       unsafe.Pointer
	arg_z       unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [6]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for reducemaxvecnorm2 kernel invocation
var reducemaxvecnorm2_args reducemaxvecnorm2_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	reducemaxvecnorm2_args.argptr[0] = unsafe.Pointer(&reducemaxvecnorm2_args.arg_x)
	reducemaxvecnorm2_args.argptr[1] = unsafe.Pointer(&reducemaxvecnorm2_args.arg_y)
	reducemaxvecnorm2_args.argptr[2] = unsafe.Pointer(&reducemaxvecnorm2_args.arg_z)
	reducemaxvecnorm2_args.argptr[3] = unsafe.Pointer(&reducemaxvecnorm2_args.arg_dst)
	reducemaxvecnorm2_args.argptr[4] = unsafe.Pointer(&reducemaxvecnorm2_args.arg_initVal)
	reducemaxvecnorm2_args.argptr[5] = unsafe.Pointer(&reducemaxvecnorm2_args.arg_n)
}

// Wrapper for reducemaxvecnorm2 CUDA kernel, asynchronous.
func k_reducemaxvecnorm2_async(x unsafe.Pointer, y unsafe.Pointer, z unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("reducemaxvecnorm2")
	}

	reducemaxvecnorm2_args.Lock()
	defer reducemaxvecnorm2_args.Unlock()

	if reducemaxvecnorm2_code == 0 {
		reducemaxvecnorm2_code = fatbinLoad(reducemaxvecnorm2_map, "reducemaxvecnorm2")
	}

	reducemaxvecnorm2_args.arg_x = x
	reducemaxvecnorm2_args.arg_y = y
	reducemaxvecnorm2_args.arg_z = z
	reducemaxvecnorm2_args.arg_dst = dst
	reducemaxvecnorm2_args.arg_initVal = initVal
	reducemaxvecnorm2_args.arg_n = n

	args := reducemaxvecnorm2_args.argptr[:]
	cu.LaunchKernel(reducemaxvecnorm2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("reducemaxvecnorm2")
	}
}

// maps compute capability on PTX code for reducemaxvecnorm2 kernel.
var reducemaxvecnorm2_map = map[int]string{0: "",
	20: reducemaxvecnorm2_ptx_20,
	30: reducemaxvecnorm2_ptx_30,
	35: reducemaxvecnorm2_ptx_35,
	50: reducemaxvecnorm2_ptx_50}

// reducemaxvecnorm2 PTX code for various compute capabilities.
const (
	reducemaxvecnorm2_ptx_20 = `
.version 4.1
.target sm_20
.address_size 64


.visible .entry reducemaxvecnorm2(
	.param .u64 reducemaxvecnorm2_param_0,
	.param .u64 reducemaxvecnorm2_param_1,
	.param .u64 reducemaxvecnorm2_param_2,
	.param .u64 reducemaxvecnorm2_param_3,
	.param .f32 reducemaxvecnorm2_param_4,
	.param .u32 reducemaxvecnorm2_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<35>;
	.reg .s64 	%rd<19>;
	// demoted variable
	.shared .align 4 .b8 reducemaxvecnorm2$__cuda_local_var_34487_32_non_const_sdata[2048];

	ld.param.u64 	%rd6, [reducemaxvecnorm2_param_0];
	ld.param.u64 	%rd7, [reducemaxvecnorm2_param_1];
	ld.param.u64 	%rd8, [reducemaxvecnorm2_param_2];
	ld.param.u64 	%rd5, [reducemaxvecnorm2_param_3];
	ld.param.f32 	%f34, [reducemaxvecnorm2_param_4];
	ld.param.u32 	%r9, [reducemaxvecnorm2_param_5];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd7;
	cvta.to.global.u64 	%rd3, %rd6;
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r11, %r16;
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd9, %r15, 4;
	add.s64 	%rd10, %rd3, %rd9;
	ld.global.f32 	%f5, [%rd10];
	add.s64 	%rd11, %rd2, %rd9;
	ld.global.f32 	%f6, [%rd11];
	mul.f32 	%f7, %f6, %f6;
	fma.rn.f32 	%f8, %f5, %f5, %f7;
	add.s64 	%rd12, %rd1, %rd9;
	ld.global.f32 	%f9, [%rd12];
	fma.rn.f32 	%f10, %f9, %f9, %f8;
	max.f32 	%f34, %f34, %f10;
	add.s32 	%r15, %r15, %r4;
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	mul.wide.s32 	%rd13, %r2, 4;
	mov.u64 	%rd14, reducemaxvecnorm2$__cuda_local_var_34487_32_non_const_sdata;
	add.s64 	%rd4, %rd14, %rd13;
	st.shared.f32 	[%rd4], %f34;
	bar.sync 	0;
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f11, [%rd4];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd15, %r12, 4;
	add.s64 	%rd17, %rd14, %rd15;
	ld.shared.f32 	%f12, [%rd17];
	max.f32 	%f13, %f11, %f12;
	st.shared.f32 	[%rd4], %f13;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f14, [%rd4];
	ld.volatile.shared.f32 	%f15, [%rd4+128];
	max.f32 	%f16, %f14, %f15;
	st.volatile.shared.f32 	[%rd4], %f16;
	ld.volatile.shared.f32 	%f17, [%rd4+64];
	ld.volatile.shared.f32 	%f18, [%rd4];
	max.f32 	%f19, %f18, %f17;
	st.volatile.shared.f32 	[%rd4], %f19;
	ld.volatile.shared.f32 	%f20, [%rd4+32];
	ld.volatile.shared.f32 	%f21, [%rd4];
	max.f32 	%f22, %f21, %f20;
	st.volatile.shared.f32 	[%rd4], %f22;
	ld.volatile.shared.f32 	%f23, [%rd4+16];
	ld.volatile.shared.f32 	%f24, [%rd4];
	max.f32 	%f25, %f24, %f23;
	st.volatile.shared.f32 	[%rd4], %f25;
	ld.volatile.shared.f32 	%f26, [%rd4+8];
	ld.volatile.shared.f32 	%f27, [%rd4];
	max.f32 	%f28, %f27, %f26;
	st.volatile.shared.f32 	[%rd4], %f28;
	ld.volatile.shared.f32 	%f29, [%rd4+4];
	ld.volatile.shared.f32 	%f30, [%rd4];
	max.f32 	%f31, %f30, %f29;
	st.volatile.shared.f32 	[%rd4], %f31;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	cvta.to.global.u64 	%rd18, %rd5;
	ld.shared.f32 	%f32, [reducemaxvecnorm2$__cuda_local_var_34487_32_non_const_sdata];
	abs.f32 	%f33, %f32;
	mov.b32 	 %r13, %f33;
	atom.global.max.s32 	%r14, [%rd18], %r13;

BB0_10:
	ret;
}


`
	reducemaxvecnorm2_ptx_30 = `
.version 4.1
.target sm_30
.address_size 64


.visible .entry reducemaxvecnorm2(
	.param .u64 reducemaxvecnorm2_param_0,
	.param .u64 reducemaxvecnorm2_param_1,
	.param .u64 reducemaxvecnorm2_param_2,
	.param .u64 reducemaxvecnorm2_param_3,
	.param .f32 reducemaxvecnorm2_param_4,
	.param .u32 reducemaxvecnorm2_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<35>;
	.reg .s64 	%rd<19>;
	// demoted variable
	.shared .align 4 .b8 reducemaxvecnorm2$__cuda_local_var_34714_32_non_const_sdata[2048];

	ld.param.u64 	%rd6, [reducemaxvecnorm2_param_0];
	ld.param.u64 	%rd7, [reducemaxvecnorm2_param_1];
	ld.param.u64 	%rd8, [reducemaxvecnorm2_param_2];
	ld.param.u64 	%rd5, [reducemaxvecnorm2_param_3];
	ld.param.f32 	%f34, [reducemaxvecnorm2_param_4];
	ld.param.u32 	%r9, [reducemaxvecnorm2_param_5];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd7;
	cvta.to.global.u64 	%rd3, %rd6;
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r11, %r16;
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd9, %r15, 4;
	add.s64 	%rd10, %rd3, %rd9;
	ld.global.f32 	%f5, [%rd10];
	add.s64 	%rd11, %rd2, %rd9;
	ld.global.f32 	%f6, [%rd11];
	mul.f32 	%f7, %f6, %f6;
	fma.rn.f32 	%f8, %f5, %f5, %f7;
	add.s64 	%rd12, %rd1, %rd9;
	ld.global.f32 	%f9, [%rd12];
	fma.rn.f32 	%f10, %f9, %f9, %f8;
	max.f32 	%f34, %f34, %f10;
	add.s32 	%r15, %r15, %r4;
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	mul.wide.s32 	%rd13, %r2, 4;
	mov.u64 	%rd14, reducemaxvecnorm2$__cuda_local_var_34714_32_non_const_sdata;
	add.s64 	%rd4, %rd14, %rd13;
	st.shared.f32 	[%rd4], %f34;
	bar.sync 	0;
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f11, [%rd4];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd15, %r12, 4;
	add.s64 	%rd17, %rd14, %rd15;
	ld.shared.f32 	%f12, [%rd17];
	max.f32 	%f13, %f11, %f12;
	st.shared.f32 	[%rd4], %f13;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f14, [%rd4];
	ld.volatile.shared.f32 	%f15, [%rd4+128];
	max.f32 	%f16, %f14, %f15;
	st.volatile.shared.f32 	[%rd4], %f16;
	ld.volatile.shared.f32 	%f17, [%rd4+64];
	ld.volatile.shared.f32 	%f18, [%rd4];
	max.f32 	%f19, %f18, %f17;
	st.volatile.shared.f32 	[%rd4], %f19;
	ld.volatile.shared.f32 	%f20, [%rd4+32];
	ld.volatile.shared.f32 	%f21, [%rd4];
	max.f32 	%f22, %f21, %f20;
	st.volatile.shared.f32 	[%rd4], %f22;
	ld.volatile.shared.f32 	%f23, [%rd4+16];
	ld.volatile.shared.f32 	%f24, [%rd4];
	max.f32 	%f25, %f24, %f23;
	st.volatile.shared.f32 	[%rd4], %f25;
	ld.volatile.shared.f32 	%f26, [%rd4+8];
	ld.volatile.shared.f32 	%f27, [%rd4];
	max.f32 	%f28, %f27, %f26;
	st.volatile.shared.f32 	[%rd4], %f28;
	ld.volatile.shared.f32 	%f29, [%rd4+4];
	ld.volatile.shared.f32 	%f30, [%rd4];
	max.f32 	%f31, %f30, %f29;
	st.volatile.shared.f32 	[%rd4], %f31;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	cvta.to.global.u64 	%rd18, %rd5;
	ld.shared.f32 	%f32, [reducemaxvecnorm2$__cuda_local_var_34714_32_non_const_sdata];
	abs.f32 	%f33, %f32;
	mov.b32 	 %r13, %f33;
	atom.global.max.s32 	%r14, [%rd18], %r13;

BB0_10:
	ret;
}


`
	reducemaxvecnorm2_ptx_35 = `
.version 4.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry reducemaxvecnorm2(
	.param .u64 reducemaxvecnorm2_param_0,
	.param .u64 reducemaxvecnorm2_param_1,
	.param .u64 reducemaxvecnorm2_param_2,
	.param .u64 reducemaxvecnorm2_param_3,
	.param .f32 reducemaxvecnorm2_param_4,
	.param .u32 reducemaxvecnorm2_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<35>;
	.reg .s64 	%rd<19>;
	// demoted variable
	.shared .align 4 .b8 reducemaxvecnorm2$__cuda_local_var_34890_32_non_const_sdata[2048];

	ld.param.u64 	%rd6, [reducemaxvecnorm2_param_0];
	ld.param.u64 	%rd7, [reducemaxvecnorm2_param_1];
	ld.param.u64 	%rd8, [reducemaxvecnorm2_param_2];
	ld.param.u64 	%rd5, [reducemaxvecnorm2_param_3];
	ld.param.f32 	%f34, [reducemaxvecnorm2_param_4];
	ld.param.u32 	%r9, [reducemaxvecnorm2_param_5];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd7;
	cvta.to.global.u64 	%rd3, %rd6;
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r11, %r16;
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB5_2;

BB5_1:
	mul.wide.s32 	%rd9, %r15, 4;
	add.s64 	%rd10, %rd3, %rd9;
	ld.global.nc.f32 	%f5, [%rd10];
	add.s64 	%rd11, %rd2, %rd9;
	ld.global.nc.f32 	%f6, [%rd11];
	mul.f32 	%f7, %f6, %f6;
	fma.rn.f32 	%f8, %f5, %f5, %f7;
	add.s64 	%rd12, %rd1, %rd9;
	ld.global.nc.f32 	%f9, [%rd12];
	fma.rn.f32 	%f10, %f9, %f9, %f8;
	max.f32 	%f34, %f34, %f10;
	add.s32 	%r15, %r15, %r4;
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB5_1;

BB5_2:
	mul.wide.s32 	%rd13, %r2, 4;
	mov.u64 	%rd14, reducemaxvecnorm2$__cuda_local_var_34890_32_non_const_sdata;
	add.s64 	%rd4, %rd14, %rd13;
	st.shared.f32 	[%rd4], %f34;
	bar.sync 	0;
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB5_6;

BB5_3:
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB5_5;

	ld.shared.f32 	%f11, [%rd4];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd15, %r12, 4;
	add.s64 	%rd17, %rd14, %rd15;
	ld.shared.f32 	%f12, [%rd17];
	max.f32 	%f13, %f11, %f12;
	st.shared.f32 	[%rd4], %f13;

BB5_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB5_3;

BB5_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB5_8;

	ld.volatile.shared.f32 	%f14, [%rd4];
	ld.volatile.shared.f32 	%f15, [%rd4+128];
	max.f32 	%f16, %f14, %f15;
	st.volatile.shared.f32 	[%rd4], %f16;
	ld.volatile.shared.f32 	%f17, [%rd4+64];
	ld.volatile.shared.f32 	%f18, [%rd4];
	max.f32 	%f19, %f18, %f17;
	st.volatile.shared.f32 	[%rd4], %f19;
	ld.volatile.shared.f32 	%f20, [%rd4+32];
	ld.volatile.shared.f32 	%f21, [%rd4];
	max.f32 	%f22, %f21, %f20;
	st.volatile.shared.f32 	[%rd4], %f22;
	ld.volatile.shared.f32 	%f23, [%rd4+16];
	ld.volatile.shared.f32 	%f24, [%rd4];
	max.f32 	%f25, %f24, %f23;
	st.volatile.shared.f32 	[%rd4], %f25;
	ld.volatile.shared.f32 	%f26, [%rd4+8];
	ld.volatile.shared.f32 	%f27, [%rd4];
	max.f32 	%f28, %f27, %f26;
	st.volatile.shared.f32 	[%rd4], %f28;
	ld.volatile.shared.f32 	%f29, [%rd4+4];
	ld.volatile.shared.f32 	%f30, [%rd4];
	max.f32 	%f31, %f30, %f29;
	st.volatile.shared.f32 	[%rd4], %f31;

BB5_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB5_10;

	cvta.to.global.u64 	%rd18, %rd5;
	ld.shared.f32 	%f32, [reducemaxvecnorm2$__cuda_local_var_34890_32_non_const_sdata];
	abs.f32 	%f33, %f32;
	mov.b32 	 %r13, %f33;
	atom.global.max.s32 	%r14, [%rd18], %r13;

BB5_10:
	ret;
}


`
	reducemaxvecnorm2_ptx_50 = `
.version 4.1
.target sm_50
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry reducemaxvecnorm2(
	.param .u64 reducemaxvecnorm2_param_0,
	.param .u64 reducemaxvecnorm2_param_1,
	.param .u64 reducemaxvecnorm2_param_2,
	.param .u64 reducemaxvecnorm2_param_3,
	.param .f32 reducemaxvecnorm2_param_4,
	.param .u32 reducemaxvecnorm2_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<35>;
	.reg .s64 	%rd<19>;
	// demoted variable
	.shared .align 4 .b8 reducemaxvecnorm2$__cuda_local_var_34890_32_non_const_sdata[2048];

	ld.param.u64 	%rd6, [reducemaxvecnorm2_param_0];
	ld.param.u64 	%rd7, [reducemaxvecnorm2_param_1];
	ld.param.u64 	%rd8, [reducemaxvecnorm2_param_2];
	ld.param.u64 	%rd5, [reducemaxvecnorm2_param_3];
	ld.param.f32 	%f34, [reducemaxvecnorm2_param_4];
	ld.param.u32 	%r9, [reducemaxvecnorm2_param_5];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd7;
	cvta.to.global.u64 	%rd3, %rd6;
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r11, %r16;
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB5_2;

BB5_1:
	mul.wide.s32 	%rd9, %r15, 4;
	add.s64 	%rd10, %rd3, %rd9;
	ld.global.nc.f32 	%f5, [%rd10];
	add.s64 	%rd11, %rd2, %rd9;
	ld.global.nc.f32 	%f6, [%rd11];
	mul.f32 	%f7, %f6, %f6;
	fma.rn.f32 	%f8, %f5, %f5, %f7;
	add.s64 	%rd12, %rd1, %rd9;
	ld.global.nc.f32 	%f9, [%rd12];
	fma.rn.f32 	%f10, %f9, %f9, %f8;
	max.f32 	%f34, %f34, %f10;
	add.s32 	%r15, %r15, %r4;
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB5_1;

BB5_2:
	mul.wide.s32 	%rd13, %r2, 4;
	mov.u64 	%rd14, reducemaxvecnorm2$__cuda_local_var_34890_32_non_const_sdata;
	add.s64 	%rd4, %rd14, %rd13;
	st.shared.f32 	[%rd4], %f34;
	bar.sync 	0;
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB5_6;

BB5_3:
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB5_5;

	ld.shared.f32 	%f11, [%rd4];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd15, %r12, 4;
	add.s64 	%rd17, %rd14, %rd15;
	ld.shared.f32 	%f12, [%rd17];
	max.f32 	%f13, %f11, %f12;
	st.shared.f32 	[%rd4], %f13;

BB5_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB5_3;

BB5_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB5_8;

	ld.volatile.shared.f32 	%f14, [%rd4];
	ld.volatile.shared.f32 	%f15, [%rd4+128];
	max.f32 	%f16, %f14, %f15;
	st.volatile.shared.f32 	[%rd4], %f16;
	ld.volatile.shared.f32 	%f17, [%rd4+64];
	ld.volatile.shared.f32 	%f18, [%rd4];
	max.f32 	%f19, %f18, %f17;
	st.volatile.shared.f32 	[%rd4], %f19;
	ld.volatile.shared.f32 	%f20, [%rd4+32];
	ld.volatile.shared.f32 	%f21, [%rd4];
	max.f32 	%f22, %f21, %f20;
	st.volatile.shared.f32 	[%rd4], %f22;
	ld.volatile.shared.f32 	%f23, [%rd4+16];
	ld.volatile.shared.f32 	%f24, [%rd4];
	max.f32 	%f25, %f24, %f23;
	st.volatile.shared.f32 	[%rd4], %f25;
	ld.volatile.shared.f32 	%f26, [%rd4+8];
	ld.volatile.shared.f32 	%f27, [%rd4];
	max.f32 	%f28, %f27, %f26;
	st.volatile.shared.f32 	[%rd4], %f28;
	ld.volatile.shared.f32 	%f29, [%rd4+4];
	ld.volatile.shared.f32 	%f30, [%rd4];
	max.f32 	%f31, %f30, %f29;
	st.volatile.shared.f32 	[%rd4], %f31;

BB5_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB5_10;

	cvta.to.global.u64 	%rd18, %rd5;
	ld.shared.f32 	%f32, [reducemaxvecnorm2$__cuda_local_var_34890_32_non_const_sdata];
	abs.f32 	%f33, %f32;
	mov.b32 	 %r13, %f33;
	atom.global.max.s32 	%r14, [%rd18], %r13;

BB5_10:
	ret;
}


`
)
